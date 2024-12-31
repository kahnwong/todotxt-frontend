<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-toolbar-title> todo.txt </q-toolbar-title>
        <q-btn flat @click="refreshPage" round dense icon="refresh" />
      </q-toolbar>
    </q-header>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
const refreshPage = () => {
  location.reload() // Reloads the current page
}
</script>
