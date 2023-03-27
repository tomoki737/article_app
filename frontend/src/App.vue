<script>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div>
    <v-app>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/article/create">作成</RouterLink>

      <RouterView />
    </v-app>
  </div>
</template>