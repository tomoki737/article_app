<template>
  <div>
    <v-card width="500" class="mx-auto mb-2" variant="outlined" >
      <v-card-item>
        <v-card-title>{{ article.title }}</v-card-title>
      </v-card-item>

      <v-card-text>
        {{ article.body }}
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
export default {
  props: {
    article: {},
  },
};
</script>