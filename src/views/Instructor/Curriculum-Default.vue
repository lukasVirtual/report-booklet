<template>
  <base-layout v-if="$route.name == 'Curriculum'">
    <v-main style="display: flex; margin: auto">
      <v-table
        style="width: 1100px; max-width: 1400px; border-radius: 10px"
        class="mt-5"
      >
        <tbody>
          <tr v-for="i in 10" :key="i">
            <td>{{ i }}</td>
            <td class="d-flex">
              <v-checkbox
                color="primary"
                class="align-center justify-center"
              ></v-checkbox>
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "CurriculumDefault",
  components: { BaseLayout },
  setup() {
    const items = [
      {
        title: "Item #1",
        value: 1,
        appendIcon: "mdi-",
      },
      {
        title: "Item #2",
        value: 2,
        appendIcon: "mdi-home",
      },
      {
        title: "Item #3",
        value: 3,
        props: {
          color: "primary",
        },
      },
    ];
    return { items };
  },
});
</script>
