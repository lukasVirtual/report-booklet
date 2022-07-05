<template>
  <v-list-item link :to="to">
    <v-list-item-icon>
      <v-tooltip right transition="slide-x-reverse-transition">
        <template #activator="{ on, attrs }">
          <v-icon v-bind="attrs" v-on="on">{{ icon }}</v-icon>
        </template>
        <span>{{ label }}</span>
      </v-tooltip>
    </v-list-item-icon>
    <v-list-item-title>{{ label }}</v-list-item-title>
  </v-list-item>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "NavigationListItem",
  props: {
    to: [String],
    label: String,
    icon: String,
  },
  setup() {
    return {};
  },
});
</script>
