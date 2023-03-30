
<template>
  <div>
    <v-container>
      <h1 class="mb-5">記事一覧</h1>
      <article-search @changeArticles="changeArticles"></article-search>
      <div v-for="(article, index) in articles" :key="index">
        <article-card
          :article="article"
          @parentGetArticles="getArticles"
        ></article-card>
      </div>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import ArticleCard from "../components/ArticleCard.vue";
import ArticleSearch from "../components/ArticleSearch.vue";
export default {
  components: {
    ArticleCard,
    ArticleSearch,
  },

  data() {
    return {
      articles: {},
    };
  },
  methods: {
    async getArticles() {
      const res = await axios.get("http://localhost:8080/articles");
      this.articles = res.data;
    },
    async changeArticles(articles) {
      this.articles = articles;
    },
  },

  mounted() {
    this.getArticles();
  },
};
</script>
