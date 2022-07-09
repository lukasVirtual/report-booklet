import { onMounted, createApp, ref } from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { loadFonts } from "./plugins/webfontloader";
import { loginService } from "./handler/loginHandler";

loadFonts();

const app = createApp(App);
try {
  const role = ref(await loginService.getUserRole());
  app.provide("role", role.value);
} catch (e) {
  console.error(e);
}
app.use(router);
app.use(vuetify);
app.mount("#app");
