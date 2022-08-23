<template>
  <div
    class="d-flex justify-space-around"
    v-if="$route.path === '/Berichtsheft'"
  >
    <v-menu
      justify-center
      ref="parentMenu"
      v-model="ruler"
      @vnode-mounted="
        timeStamp ? (input = timeStamp) : input;
        idx = index;
      "
    >
      <template v-slot:activator="{ props }">
        <v-btn
          @click="ruler = true"
          color="blue"
          v-bind="props"
          v-model="input"
          v-text="input"
          style="margin-bottom: 30px"
        >
        </v-btn>
      </template>
      <v-list height="150">
        <v-list-item
          v-for="(item, index) in times"
          :key="index"
          :value="index"
          @click="(input = item.title) && (ruler = false)"
          @blur="$emit('updateTime')"
        >
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script lang="ts">
import { store } from "@/handler/store";
import { defineComponent, ref, watchEffect } from "vue";

export default defineComponent({
  name: "TimerTextField",
  props: {
    timeStamp: String,
    index: Number,
  },
  emits: ["updateTime"],

  setup(props, { emit }) {
    const times = [
      { title: "00:00" },
      { title: "00:30" },
      { title: "01:30" },
      { title: "02:00" },
      { title: "02:30" },
      { title: "03:00" },
      { title: "03:30" },
      { title: "04:00" },
      { title: "04:30" },
      { title: "05:00" },
      { title: "05:30" },
      { title: "06:00" },
      { title: "06:30" },
      { title: "07:00" },
      { title: "07:30" },
      { title: "08:00" },
      { title: "08:30" },
      { title: "09:00" },
      { title: "09:30" },
      { title: "10:00" },
    ];
    let input = ref<string | undefined>("00:00");
    let ruler = ref(false);
    let idx = ref<number | undefined>(undefined);

    watchEffect(() => {
      store.time = input.value as string;
      store.input = "";
      store.id = idx.value;
    });

    return {
      times,
      input,
      ruler,
      idx,
    };
  },
});
</script>
