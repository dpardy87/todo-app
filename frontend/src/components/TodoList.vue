<template>
  <div class="flex w-full"> <!-- full width container -->
    <div class="w-1/2 p-4"> <!-- padding of val 4: 1 rem or 16px -->
      <TodoForm @add="addTodo"></TodoForm>
    </div>
    <div class="w-1/2 p-4">
      <ul>
        <!-- Only show tasks that are not completed -->
        <li v-for="todo in activeTodos" :key="todo.id" class="bg-white rounded-lg shadow p-3 mb-4">
          <div class="flex justify-between items-center">
            <div>
              <h3 class="text-lg font-semibold">{{ todo.task }}</h3>
              <p class="text-gray-600">{{ todo.description }}</p>
            </div>
            <div>
              <input v-model="todo.completed" type="checkbox" class="mr-2" @change="updateTodoStatus(todo)">
              <span class="text-sm">{{ todo.completed ? 'Completed' : 'Incomplete' }}</span>
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
  components: {
    TodoForm
  },
  data() {
    return {
      todos: [],
    };
  },
  computed: {
    // Filter out completed todos
    activeTodos() {
      return this.todos.filter(todo => !todo.completed);
    }
  },
  mounted() {
    this.fetchTodos();
  },
  methods: {
    async fetchTodos() {
      try {
        const response = await axios.get('http://localhost:8080/api/todos');
        this.todos = response.data;
      } catch (error) {
        console.error("There was an error!", error);
      }
    },
    async addTodo(newTodo) {
      try {
        const response = await axios.post('http://localhost:8080/api/todos', {
          task: newTodo.taskName,
          description: newTodo.description,
          completed: false // New todos are always incomplete
        });
        this.todos.push(response.data);
      } catch (error) {
        console.error("There was an error when adding a todo", error);
      }
    },
    async updateTodoStatus(todo) {
      try {
        await axios.put(`/api/todos/${todo.id}`, { completed: todo.completed });
      } catch (error) {
        console.error("There was an error updating the todo status", error);
      }
    }
  }
};
</script>
