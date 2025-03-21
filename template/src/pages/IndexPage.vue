<template>
  <q-page class="row justify-evenly">
    <div class="q-pa-md example-row-equal-width">
      <div class="row">
        <example-component
          title="Example component"
          active
          :todos="todosFromApi"
        ></example-component>

        <p>{{ todosFromApi }}</p>
      </div>

      <div class="row">
        <example-component title="Example component" active :todos="todos"></example-component>
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Todo } from 'components/models'
import ExampleComponent from 'components/ExampleComponent.vue'
import axios from 'axios'

const todos = ref<Todo[]>([
  {
    id: 1,
    project: 'fo',
    todo: 'bar',
    context: 'ct1',
  },
  {
    id: 2,
    project: 'fo',
    todo: 'bar',
    context: 'ct1',
  },
])

// api
const todosFromApi = ref<Todo[]>([])

const fetchData = async () => {
  const response = await axios.get('/api/todo/today')
  todosFromApi.value = response.data as Todo[]
}

// Fetch data when component is mounted
onMounted(fetchData)
</script>
