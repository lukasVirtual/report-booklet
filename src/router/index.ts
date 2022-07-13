import { createRouter, createWebHistory } from "vue-router";
import App from "../App.vue";
import LoginDefault from "../Auth/Login-Default.vue";
import RegisterDefault from "../Auth/Register-Default.vue";
import MainPage from "../views/Berichtsheft/Main.vue";
import DashboardDefault from "../views/Dashboard/Dashboard-Default.vue";
import AdminDefault from "../views/Admin/Admin-Default.vue";
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
        { path: "/Dashboard", name: "Dashboard", component: DashboardDefault, meta: { requiresAuth: true} },
        { path: "/Berichtsheft", name: "Berichtsheft", component: MainPage, meta: { requiresAuth: true } },
        { path: "/Admin", name: "Admin", component: AdminDefault, meta: { requiresAuth: true } },
      ],
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

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth) {
    try {
      const statusCheck = await loginService.checkStatus();
      console.log(statusCheck);
      if (statusCheck) next();
      else next({ path: "/Login" })
    } catch (e) {
      next({ path: "/Login" });
    }
  } else {
    next()
  }

})
export default router;
