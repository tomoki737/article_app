<template>
  <div>
    <v-card width="500" class="mx-auto mb-2" variant="outlined">
      <v-row>
        <v-card-item>
          <router-link
            :to="{
              name: 'articles.edit',
              params: { id: article.id },
            }"
            >編集</router-link
          >
          <v-card-title>{{ article.title }}</v-card-title>
        </v-card-item>
        <v-card-item>
          <v-btn variant="outlined" @click="deleteArticle">削除</v-btn>
        </v-card-item>
      </v-row>

      <v-card-text>
        {{ article.body }}
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  props: {
    article: {},
  },
  methods: {
    async deleteArticle() {
      const res = await axios.delete("http://localhost:8080/articles/" + this.article.id)
    }
  }
};
</script>