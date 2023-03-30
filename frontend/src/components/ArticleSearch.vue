<template>
  <div>
    <v-row justify="end" class="mb-5">
      <v-col cols="5">
        <v-text-field
          label="Title"
          variant="outlined"
          v-model="title"
        ></v-text-field>
      </v-col>
      <v-col cols="5">
        <v-text-field
          label="Body"
          variant="outlined"
          v-model="body"
        ></v-text-field>
      </v-col>
      <v-col cols="2">
        <v-btn @click="searchArticle" class="me-0" color="blue" style="height: 56px">
          検索
        </v-btn>
      </v-col>
    </v-row>
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
    async searchArticle() {
      const res = await axios
        .delete(
          "http://localhost:8080/articles/search?title=" +
            this.title +
            "&body=" +
            this.body
        )
        .catch((err) => {
          console.error(err);
          return;
        });
      this.$emit("changeArticles", res.data);
    },
  },
};
</script>