<template>
  <q-page class="row justify-evenly">
    <div class="q-pa-md example-row-equal-width">
      <div class="row">
        <example-component
          title="Example component"
          active
          :todos="todos"
          :meta="meta"
        ></example-component>

        <p>{{ todosFromApi }}</p>
      </div>

      <div class="row">
        <example-component
          title="Example component"
          active
          :todos="todos"
          :meta="meta"
        ></example-component>
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Todo, Todo2, Meta } from 'components/models'
import ExampleComponent from 'components/ExampleComponent.vue'
import axios from 'axios'

const todos = ref<Todo[]>([
  {
    id: 1,
    content: 'ct1',
  },
  {
    id: 2,
    content: 'ct2',
  },
  {
    id: 3,
    content: 'ct3',
  },
  {
    id: 4,
    content: 'ct4',
  },
  {
    id: 5,
    content: 'ct5',
  },
])

const meta = ref<Meta>({
  totalCount: 1200,
})

// api
const todosFromApi = ref<Todo2[]>([])

const fetchData = async () => {
  const response = await axios.get('/api/todo/today')
  todosFromApi.value = response.data as Todo2[]
}

// Fetch data when component is mounted
onMounted(fetchData)
</script>
