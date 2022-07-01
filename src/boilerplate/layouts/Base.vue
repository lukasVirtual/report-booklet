<template>
  <v-app
    :theme="theme"
    style="display: flex; align-items: center; justify-content: center"
  >
    <v-app-bar :color="theme" style="position: fixed">
      <v-spacer></v-spacer>
      <div class="text-right" style="margin: 20px">
        <v-btn rounded @click="switchTheme"
          ><v-icon :icon="iconValue"></v-icon
        ></v-btn>
        <v-btn rounded @click="logout"
          ><v-icon style="width: 35px">mdi-account-arrow-left</v-icon
          >Logout</v-btn
        >
      </div></v-app-bar
    >
    <v-navigation-drawer>
      <v-list-item height="70px" :to="{ path: '/Berichtsheft' }"
        ><v-icon style="width: 40px">mdi-file-document-edit-outline</v-icon>
        Berichtsheft</v-list-item
      >
      <v-list-item height="70px" :to="{ path: '/Dashboard' }"
        ><v-icon style="width: 40px">mdi-monitor-dashboard</v-icon>
        Dashboard</v-list-item
      >
    </v-navigation-drawer>
    <Suspense>
      <slot></slot>
      <template #fallback>
        <v-overlay v-model="overlay"> </v-overlay>
      </template>
    </Suspense>
  </v-app>
</template>
<script lang="ts">
import { defineComponent, ref } from "vue";
import { useRoute } from "vue-router";

export default defineComponent({
  name: "BaseLayout",

  data() {
    let isAuthenticated = ref(false);
    let theme = ref("dark");
    const router = useRoute();
    const overlay = ref(true);
    let iconValue =
      theme.value == "dark"
        ? ref("mdi-moon-waning-crescent")
        : ref("mdi-weather-sunny");

    const logout = () => {
      localStorage.removeItem("authenticated");
      localStorage.setItem("authenticated", "false");
      this.$router.push({ name: "Login", query: { redirect: "/Login" } });
    };

    const switchTheme = () => {
      theme.value = theme.value == "dark" ? "light" : "dark";
      iconValue.value =
        theme.value == "dark"
          ? "mdi-moon-waning-crescent"
          : "mdi-weather-sunny";
    };

    return { isAuthenticated, logout, theme, overlay, switchTheme, iconValue };
  },
});
</script>