<template>
  <v-app :theme="theme" style="display: flex; align-items: center; justify-content: center">
    <v-app-bar :color="theme" style="position: fixed">
      <v-btn @click="togglelistState" icon><v-icon style="width: 35px"
          size="22">mdi-format-list-bulleted</v-icon></v-btn>

      <slot name="navIcons"></slot>

      <v-spacer></v-spacer>
      <div class="text-right" style="margin: 20px">
        <v-btn rounded v-if="showHearbeat"><v-icon color="blue" style="width: 35px"
            size="22">mdi-access-point</v-icon></v-btn>
        <v-btn rounded @click="logout"><v-icon style="width: 35px">mdi-account-arrow-left</v-icon>Logout</v-btn>
        <v-btn rounded @click="switchTheme"><v-icon :icon="iconValue"></v-icon></v-btn>
      </div>
      <v-img class="logo--light" max-width="80" lazy-src="src\assets\vhf-logo.png"
        src="src\assets\vhf-logo.png"></v-img>
    </v-app-bar>
    <v-navigation-drawer :rail="toggleList">
      <v-list nav>
        <v-list-item v-if="role === 'user'" :to="{ name: 'Berichtsheft' }" height="50"
          prepend-icon="mdi-file-document-edit-outline" title="Berichtsheft" value="Berichtsheft"></v-list-item>

        <v-list-item v-if="role === 'user'" :to="{ name: 'Dashboard' }" height="50" prepend-icon="mdi-monitor-dashboard"
          title="Dashboard" value="Dashboard"></v-list-item>

        <v-list-item v-if="role === 'instructor'" :to="{ name: 'Instructor' }" height="50"
          prepend-icon="mdi-account-search-outline" title="Control Center" value="Control Center"></v-list-item>

        <v-list-item v-if="role === 'instructor'" :to="{ name: 'Curriculum' }" height="50"
          prepend-icon="mdi-beaker-check" title="Curriculum" value="Curriculum"></v-list-item>

        <v-list-item v-if="role === 'admin'" :to="{ name: 'Admin' }" height="50"
          prepend-icon="mdi-account-multiple-plus-outline" title="Admin Center" value="Admin Center"></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <Suspense>
      <template #default>
        <slot></slot>
      </template>
      <template #fallback>
        <v-overlay v-model="overlay"></v-overlay>
      </template>
    </Suspense>
    <!-- <v-overlay :model-value="loading" class="align-center justify-center">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay> -->
  </v-app>
</template>
<script lang="ts">
// import type { LoginServiceInterface } from "@/handler/loginHandler";
import { loginService } from "@/handler/loginHandler"
import { io } from "socket.io-client";
import { defineComponent, onMounted, ref } from "vue";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "BaseLayout",

  setup() {
    // const loginService = inject(
    //   "provide-login-service"
    // ) as LoginServiceInterface;
    let isAuthenticated = ref(false);
    const loading = ref<undefined | boolean>(undefined);
    let theme = ref("dark");
    const router = useRouter();
    const overlay = ref(true);
    const showHearbeat = ref(false);
    const toggleList =
      localStorage.getItem("list-state") == "true" ? ref(true) : ref(false);
    const socket = io("http://localhost:7000");

    console.log("value: ", toggleList.value);
    let role = ref("");
    onMounted(async () => {
      role.value = (await loginService.getUserRole()) as unknown as string;
    });
    let iconValue =
      theme.value == "dark"
        ? ref("mdi-moon-waning-crescent")
        : ref("mdi-weather-sunny");

    const logout = async () => {
      // loading.value = true;
      socket.emit("logout", {
        user_name: await loginService.getUser(),
        user_id: await loginService.getUserID(),
      });
      const logoutHandling = await loginService.logout();
      console.log(logoutHandling);

      // loading.value = false;
      router.push({ path: "/Login" });
    };

    const switchTheme = () => {
      theme.value = theme.value == "dark" ? "light" : "dark";
      iconValue.value =
        theme.value == "dark"
          ? "mdi-moon-waning-crescent"
          : "mdi-weather-sunny";
    };

    const togglelistState = () => {
      localStorage.removeItem("list-state");
      toggleList.value = !toggleList.value;
      localStorage.setItem("list-state", `${toggleList.value}`);
    };

    setInterval(() => {
      showHearbeat.value = !showHearbeat.value;
    }, 400);

    showHearbeat.value = false;

    return {
      isAuthenticated,
      showHearbeat,
      toggleList,
      theme,
      overlay,
      iconValue,
      role,
      loading,
      togglelistState,
      logout,
      switchTheme,
    };
  },
});
</script>
