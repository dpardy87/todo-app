<template>
  <div class="p-4 bg-gray-100">
    <form @submit.prevent="addTodo">
      <div class="mb-4"><!-- margin-bottom, 1rem or 16px -->
        <label for="taskName" class="block text-gray-700 font-bold mb-1">Task Name:</label>
        <input type="text" id="taskName" v-model="taskName" class="w-full p-2 border rounded-md focus:outline-none focus:border-blue-500"><!-- w-full = width 100% of its parent, p-2 = 8px -->
      </div>
      <div class="mb-4">
        <label for="description" class="block text-gray-700 font-bold mb-1">Description:</label>
        <textarea id="description" v-model="description" class="w-full p-4 border rounded-md focus:outline-none focus:border-blue-500" rows="4"></textarea>
      </div>
      <div class="flex mb-4"><!-- flex: makes child elements lay out in a row. default is horizontal -->
        <div class="flex h-5">
            <input type="checkbox" v-model="completed" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
        </div>
        <div class="ms-2 text-sm lift-text">
            <label for="helper-checkbox" class="font-medium">Completed</label>
            <p id="helper-checkbox-text" class="text-xs font-normal dark:text-gray-300">Check the checkbox to know it's completed</p>
        </div>
      </div>
      <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Add Task</button><!-- py-val px-val: adds padding -->
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      taskName: '',
      description: '',
      completed: false
    };
  },
  methods: {
    addTodo() {
      // emit: notify parent component about new todo
      this.$emit('add', {
        taskName: this.taskName,
        description: this.description,
        completed: this.completed
      });
      // reset fields
      this.taskName = '';
      this.description = '';
      this.completed = false;
    }
  }
};
</script>

<style scoped>
  .lift-text {
    transform: translateY(-2px);
  }
</style>
