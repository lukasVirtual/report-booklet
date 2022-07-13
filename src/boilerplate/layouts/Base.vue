<template>
  <v-app
    :theme="theme"
    style="display: flex; align-items: center; justify-content: center"
  >
    <v-app-bar :color="theme" style="position: fixed">
      <v-btn @click="togglelistState" icon
        ><v-icon style="width: 35px" size="22"
          >mdi-format-list-bulleted</v-icon
        ></v-btn
      >
      <v-spacer></v-spacer>
      <div class="text-right" style="margin: 20px">
        <v-btn rounded v-if="showHearbeat"
          ><v-icon color="blue" style="width: 35px" size="22"
            >mdi-access-point</v-icon
          ></v-btn
        >
        <v-btn rounded @click="logout"
          ><v-icon style="width: 35px">mdi-account-arrow-left</v-icon
          >Logout</v-btn
        >
        <v-btn rounded @click="switchTheme"
          ><v-icon :icon="iconValue"></v-icon
        ></v-btn>
      </div>
      <v-img
        class="logo--light"
        max-width="80"
        src="src\assets\vhf-logo.png"
      ></v-img>
    </v-app-bar>
    <v-navigation-drawer v-model="toggleList">
      <v-list-item
        v-if="role === 'user'"
        height="60px"
        :to="{ name: 'Berichtsheft' }"
        ><v-icon style="width: 40px">mdi-file-document-edit-outline</v-icon>
        Berichtsheft</v-list-item
      >
      <v-list-item
        v-if="role === 'user'"
        height="60px"
        :to="{ name: 'Dashboard' }"
        ><v-icon style="width: 40px">mdi-monitor-dashboard</v-icon>
        Dashboard</v-list-item
      >
      <v-list-item
        v-if="role === 'instructor'"
        height="60px"
        :to="{ name: 'Dashboard' }"
        ><v-icon style="width: 40px">mdi-monitor</v-icon> Control
        Center</v-list-item
      >
      <v-list-item v-if="role === 'admin'" height="60px" :to="{ name: 'Admin' }"
        ><v-icon style="width: 40px">mdi-account-multiple-plus-outline</v-icon>
        Admin Center</v-list-item
      >
    </v-navigation-drawer>
    <Suspense>
      <template #default>
        <slot></slot>
      </template>
      <template #fallback>
        <v-overlay v-model="overlay"></v-overlay>
      </template>
    </Suspense>
  </v-app>
</template>
<script lang="ts">
import { loginService } from "@/handler/loginHandler";
import { defineComponent, inject, onMounted, onServerPrefetch, ref } from "vue";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "BaseLayout",
  // components: { NavigationListItem },
  setup() {
    let isAuthenticated = ref(false);
    let theme = ref("dark");
    const router = useRouter();
    const overlay = ref(true);
    const showHearbeat = ref(false);
    const toggleList = ref(true);
    let role = ref("");
    onMounted(() => {
      role.value = inject("role") as string;
    });
    let iconValue =
      theme.value == "dark"
        ? ref("mdi-moon-waning-crescent")
        : ref("mdi-weather-sunny");

    const logout = async () => {
      const logoutHandling = await loginService.logout();
      console.log(logoutHandling);

      router.push({ name: "Login", query: { redirect: "/Login" } });
    };

    const switchTheme = () => {
      theme.value = theme.value == "dark" ? "light" : "dark";
      iconValue.value =
        theme.value == "dark"
          ? "mdi-moon-waning-crescent"
          : "mdi-weather-sunny";
    };

    const togglelistState = () => {
      toggleList.value = !toggleList.value;
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
      togglelistState,
      logout,
      switchTheme,
    };
  },
});
</script>
