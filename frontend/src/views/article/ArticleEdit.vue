<template>
  <div>
    <v-container>
      <h1 class="mb-10">記事作成</h1>
      <div style="max-width: 700px" class="mx-auto">
        <v-text-field v-model="articleForm.title" label="タイトル"></v-text-field>
        <v-textarea v-model="articleForm.body" label="本文"></v-textarea>
        <v-btn @click="storeArticle" block class="mt-3" color="blue">
          編集
        </v-btn>
      </div>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
export default {
  props: {
    id: "",
  },

  data() {
    return {
      articleForm: {
        title: "",
        body: "",
      },
    };
  },

  methods: {
    async storeArticle() {
      const res = await axios.post("http://localhost:8080/articles", this.articleForm);

      this.$router.push("/");
    },

    async getArticle() {
      const res = await axios.get(
        "http://localhost:8080/articles/" + this.id
      );
      this.articleForm.title = res.data.title
      this.articleForm.body = res.data.body
    },
  },
  mounted() {
    this.getArticle();
  },
};
</script>