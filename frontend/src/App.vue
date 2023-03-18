<script>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div>
    <v-app>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>

      <RouterView />
    </v-app>
  </div>
</template>