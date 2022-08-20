import { createApp, ref } from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { loadFonts } from "./plugins/webfontloader";
import { loginService } from "./handler/loginHandler";
import Toast, { POSITION } from "vue-toastification";
import "vue-toastification/dist/index.css";

loadFonts();

const app = createApp(App);
export async function boot(): Promise<void> {
  try {
    const role = ref(await loginService.getUserRole());
    app.provide("role", role.value);
  } catch (e) {
    console.error(e);
  }
}
app.use(router);
app.use(vuetify);
app.use(Toast, {
  position: POSITION.TOP_CENTER,
  timeout: 2000,
});
app.mount("#app");
