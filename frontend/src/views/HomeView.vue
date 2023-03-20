
<template>
  <div>
    <v-container>
      <h1 class="mb-5">記事一覧</h1>
      <v-btn icon="mdi-home" />
      <div v-for="(article, index) in articles" :key="index">
        <article-card :article="article"></article-card>
      </div>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import ArticleCard from "../components/ArticleCard.vue";
export default {
  components: {
    ArticleCard,
  },

  data() {
    return {
      articles: {},
    };
  },
  methods: {
    async getArticles() {
      const res = await axios.get("http://localhost:8080/");
      this.articles = res.data;
      console.log(this.articles);
    },
  },

  mounted() {
    this.getArticles();
  },
};
</script>
