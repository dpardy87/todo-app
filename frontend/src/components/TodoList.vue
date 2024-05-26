<template>
  <div class="flex w-full"> <!-- full width container -->
    <div class="w-1/2 p-4"> <!-- padding of val 4: 1 rem or 16px -->
      <TodoForm @add="addTodo"></TodoForm>
    </div>
    <div class="w-1/2 p-4">
      <ul>
        <li v-for="todo in todos" :key="todo.id" class="bg-white rounded-lg shadow p-3 mb-4">
          <div class="flex justify-between items-center">
            <div>
              <h3 class="text-lg font-semibold">{{ todo.task }}</h3>
              <p class="text-gray-600">{{ todo.description }}</p>
            </div>
            <div>
              <input type="checkbox" v-model="todo.completed" class="mr-2">
              <span class="text-sm">{{ todo.completed ? 'Completed' : 'Incomplete' }}</span>
              <button @click="deleteTodo(todo.id)" class="ml-2 text-red-500 hover:text-red-700">Delete</button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>


<script>
import TodoForm from './TodoForm.vue';
import axios from 'axios';

export default {
  data() {
    return {
      todos: [],
    };
  },
  components: {
    TodoForm
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
    async addTodo(newTodo) {
      try {
        console.log("newTodo:", newTodo)
        const response = await axios.post('/api/todos', { task: newTodo.taskName, description: newTodo.description, completed: newTodo.completed });
        console.log("adding a todo, response constant:", response)
        this.todos.push(response.data);
      } catch (error) {
        console.error("There was an error when adding a todo", error);
      }
    },
    async deleteTodo(id) {
      try {
        await axios.delete(`/api/todos/${id}`);
        this.todos = this.todos.filter(todo => todo.id !== id);
        console.log(`Successfully deleted todo with ID ${id}`)
      } catch (error) {
        console.error(`There was an error deleting the todo ${id}`, error);
      }
    }
  },
  mounted() {
    this.fetchTodos();
  }
};
</script>
