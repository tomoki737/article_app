<template>
  <div>
    <v-container>
      <h1 class="mb-10">記事作成</h1>
      <div style="max-width: 700px" class="mx-auto">
        <v-text-field v-model="title" label="タイトル"></v-text-field>
        <v-textarea v-model="body" label="本文"></v-textarea>
        <v-btn @click="storeArticle" block class="mt-3" color="blue">
          投稿
        </v-btn>
      </div>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      title: "",
      body: "",
    };
  },

  methods: {
    async storeArticle() {
      const articleParams = new URLSearchParams();
      articleParams.append("title", this.title);
      articleParams.append("body", this.body);

      const res = await axios.post("http://localhost:8080/articles", articleParams);

      this.$router.push("/");
    },
  },
};
</script>