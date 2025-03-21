<template>
  <div class="q-pa-md" style="max-width: 400px">
    <q-list bordered separator>
      <q-item v-ripple v-for="todo in todos" :key="todo.id">
        <q-item-section>
          <p>
            <span class="todo-context">{{ todo.context }}</span>
            <span class="todo-project">{{ todo.project }}</span>
            <span> {{ todo.todo }} </span>
          </p>
        </q-item-section>
      </q-item>
    </q-list>
  </div>
</template>

<script setup lang="ts">
import type { Todo } from 'components/models'

export interface TodoProps {
  todos: Todo[]
}

withDefaults(defineProps<TodoProps>(), {})
</script>
