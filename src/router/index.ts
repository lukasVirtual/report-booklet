import { createRouter, createWebHistory } from "vue-router";
// import TextField from "../boilerplate/Text-Field.vue"
import App from "../App.vue";
import LoginDefault from "../Login/Login-Default.vue";
import RegisterDefault from "../Login/Register-Default.vue";
import MainPage from "../views/Berichtsheft/Main.vue";
import DashboardDefault from "../views/Dashboard/Dashboard-Default.vue";
import BaseLayout from "../boilerplate/layouts/Base.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      redirect: "/Berichtsheft",
      name: "home",
      component: App,
      children: [
        { path: "/Dashboard", name: "Dashboard", component: DashboardDefault },
        { path: "/Berichtsheft", name: "Berichtsheft", component: MainPage },
      ],
      beforeEnter: async (to, from, next) => {
        const isAuthenticated =
          localStorage.getItem("authenticated") === "true" ? true : false;
        console.log(to.path);
        if (to.path !== "/Login" && !isAuthenticated) {
          next({ path: "/Login" });
        } else {
          next();
        }
      },
    },
    {
      path: "/Login",
      name: "Login",
      component: LoginDefault,
    },
    {
      path: "/Register",
      name: "Register",
      component: RegisterDefault,
    },
  ],
});
export default router;
