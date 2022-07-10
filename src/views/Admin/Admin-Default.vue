<template>
  <base-layout v-if="$route.name === 'Admin'">
    <v-main style="display: flex; margin: auto">
      <h2 class="font-weight-bold text-center mt-5">Admin Center</h2>
      <v-table style="width: 1200px" class="mt-5">
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Role</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(elem, idx) in resultArr" :key="idx">
            <td>{{ elem.name }}</td>
            <td>{{ elem.role }}</td>
          </tr>
        </tbody>
      </v-table>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { loginService } from "@/handler/loginHandler";
import { defineComponent, onMounted } from "@vue/runtime-core";
import { ref } from "vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "AdminDefault",
  components: { BaseLayout },

  setup() {
    const resultArr = ref<[{ name: string; role: string }]>([
      { name: "", role: "" },
    ]);

    onMounted(async () => {
      const data = await loginService.getAllData();
      for (const elem of data) {
        resultArr.value.push({
          name: elem.Name,
          role: elem.Role,
        });
      }

      resultArr.value.map((value, idx) => {
        if (value.name == "") resultArr.value.splice(idx, 1);
      });
    });
    return {
      resultArr,
    };
  },
});
</script>