<template>
  <q-page class="row justify-evenly">
    <div class="q-pa-md example-row-equal-width">
      <div class="row">
        <todo-component title="Today" active :todos="todoToday"></todo-component>
      </div>

      <div class="row">
        <todo-component title="Tinkering" active :todos="todoTinkering"></todo-component>
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Todo } from 'components/models'
import TodoComponent from 'components/TodoComponent.vue'
import axios from 'axios'

const todoToday = ref<Todo[]>([])
const todoTinkering = ref<Todo[]>([])

const fetchTodoToday = async () => {
  const response = await axios.get('/api/todo/today')
  todoToday.value = response.data as Todo[]
}

const fetchTodoTinkering = async () => {
  const response = await axios.get('/api/todo/tinkering')
  todoTinkering.value = response.data as Todo[]
}

// Fetch data when component is mounted
onMounted(fetchTodoToday)
onMounted(fetchTodoTinkering)
</script>
