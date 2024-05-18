<template>
  <div>
    <input v-model="newTodo" @keyup.enter="addTodo">
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        {{ todo.task }}
        <button @click="deleteTodo(todo.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      todos: [],
      newTodo: ''
    };
  },
  methods: {
    async fetchTodos() {
      try {
        const response = await axios.get('/api/todos');
        this.todos = response.data;
      } catch (error) {
        console.error("There was an error!", error);
      }
    },
    async addTodo() {
      try {
        const response = await axios.post('/api/todos', { task: this.newTodo, completed: false });
        console.log("adding todo, response constant:", response)
        this.todos.push(response.data);
        this.newTodo = '';
      } catch (error) {
        console.error("There was an error wheb adding a todo", error);
      }
    },
    async deleteTodo(id) {
      try {
        await axios.delete(`/api/todos/${id}`);
        this.todos = this.todos.filter(todo => todo.id !== id);
      } catch (error) {
        console.error("There was an error!", error);
      }
    }
  },
  mounted() {
    this.fetchTodos();
  }
};
</script>
