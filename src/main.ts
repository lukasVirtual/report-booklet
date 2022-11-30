import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { loadFonts } from "./plugins/webfontloader";
import Toast, { POSITION } from "vue-toastification";
import { createPinia } from "pinia";
import "vue-toastification/dist/index.css";
import { ProvideDataService } from "./handler/dataHandler";

loadFonts();

const pinia = createPinia();
const app = createApp(App);

app.use(router);
app.use(vuetify);
app.use(pinia);
app.use(Toast, {
  position: POSITION.TOP_CENTER,
  timeout: 1500,
});
// app.provide("provide-login-service", ProvideLoginService());
app.provide("provide-data-service", ProvideDataService());
app.mount("#app");
