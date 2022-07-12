<template>
  <div v-if="$route.path == '/Berichtsheft'">
    <v-menu anchor="center" transition="slide-x-transition">
      <template v-slot:activator="{ props }">
        <v-btn
          v-bind="props"
          style="margin-bottom: 30px; border-radius: 15px; min-width: 20px"
          elevation="20"
          ><v-icon size="25">mdi-text-box-outline</v-icon></v-btn
        >
      </template>
      <v-list width="400px" height="500">
        <v-list-item v-for="num in 16" :key="num">
          <v-list-item-title>
            <v-list-item-title>The {{ num }} Qualification</v-list-item-title>
          </v-list-item-title>
          <v-spacer></v-spacer>
          <v-list-item-action>
            <v-checkbox color="deep-purple accent-4"></v-checkbox>
          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "QualificationsDefault",

  setup() {
    return {};
  },
});
</script>
