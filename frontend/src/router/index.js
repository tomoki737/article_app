import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ArticleCreate from "../views/article/ArticleCreate.vue";
import ArticleEdit from "../views/article/ArticleEdit.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/article/create",
      name: "article.create",
      component: ArticleCreate,
    },
    {
      path: "/article/:id/edit",
      name: "articles.edit",
      component: ArticleEdit,
      props: true
    },
  ],
});

export default router;
