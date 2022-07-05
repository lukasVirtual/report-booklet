<template>
  <v-app
    :theme="theme"
    style="display: flex; align-items: center; justify-content: center"
  >
    <v-navigation-drawer v-model="toggleList">
      <v-list-item height="70px" :to="{ name: 'Berichtsheft' }"
        ><v-icon style="width: 40px">mdi-file-document-edit-outline</v-icon>
        Berichtsheft</v-list-item
      >
      <v-list-item link height="70px" :to="{ name: 'Dashboard' }"
        ><v-icon style="width: 40px">mdi-monitor-dashboard</v-icon>
        Dashboard</v-list-item
      >
      <!-- <navigation-list-item
        icon="mdi-monitor-dashboard"
        :to="{ name: 'Dashboard' }"
      ></navigation-list-item> -->
    </v-navigation-drawer>
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
import NavigationListItem from "../NavigationListItem.vue";

export default defineComponent({
  name: "BaseLayout",
  components: { NavigationListItem },
  data() {
    let isAuthenticated = ref(false);
    let theme = ref("dark");
    // const router = useRoute();
    const overlay = ref(true);
    const showHearbeat = ref(false);
    const toggleList = ref(true);
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

    const togglelistState = () => {
      toggleList.value = !toggleList.value;
    };

    setInterval(() => {
      showHearbeat.value = !showHearbeat.value;
    }, 500);

    showHearbeat.value = false;

    return {
      isAuthenticated,
      showHearbeat,
      toggleList,
      theme,
      overlay,
      iconValue,
      togglelistState,
      logout,
      switchTheme,
    };
  },
});
</script>
