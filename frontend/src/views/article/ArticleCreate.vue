<template>
  <div>
    <v-container>
      <h1 class="mb-10">記事作成</h1>
      <div style="max-width: 700px" class="mx-auto">
        <v-text-field label="タイトル"></v-text-field>
        <v-textarea label="本文"></v-textarea>
        <v-btn block class="mt-3" color="blue">投稿</v-btn>
      </div>
    </v-container>
  </div>
</template>

<script>
export default {};
</script>