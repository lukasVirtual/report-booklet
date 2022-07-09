import { createRouter, createWebHistory } from "vue-router";
import App from "../App.vue";
import LoginDefault from "../Login/Login-Default.vue";
import RegisterDefault from "../Login/Register-Default.vue";
import MainPage from "../views/Berichtsheft/Main.vue";
import DashboardDefault from "../views/Dashboard/Dashboard-Default.vue";
import { loginService } from "@/handler/loginHandler";

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
        try {
          const statusCheck = await loginService.checkStatus();
          console.log(statusCheck);
          if (statusCheck) next();
        } catch (e) {
          next({ path: "/Login" });
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
