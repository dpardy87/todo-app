<template>
  <div class="p-4 bg-gray-100">
    <form @submit.prevent="addTodo">
      <div class="mb-4"><!-- margin-bottom, 1rem or 16px -->
        <label for="taskName" class="block text-gray-700 font-bold mb-1">Task Name:</label>
        <input id="taskName" ref="taskInput" v-model="taskName" type="text" class="w-full p-2 border rounded-md focus:outline-none focus:border-blue-500"><!-- w-full = width 100% of its parent, p-2 = 8px -->
      </div>
      <div class="mb-4">
        <label class="block text-gray-700 font-bold mb-1">Description:</label>
        <textarea v-model="description" class="w-full p-4 border rounded-md focus:outline-none focus:border-blue-500" rows="4"></textarea>
      </div>
      <div class="flex mb-4"><!-- flex: makes child elements lay out in a row. default is horizontal -->
        <div class="ms-2 text-sm lift-text">
            <p class="font-normal dark:text-gray-300">Tasks Default to Incomplete</p>
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
  mounted() {
    this.$refs.taskInput.focus();
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
