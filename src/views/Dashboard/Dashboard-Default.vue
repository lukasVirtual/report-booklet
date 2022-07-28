<template v-if="$route.path === '/Dashboard'">
  <base-layout>
    <v-main style="margin: 1.5rem">
      <v-card height="100%" width="100%">
        <div
          class="text-center d-flex align-center justify-center mb-6"
          style="height: 75%"
        >
          <v-progress-circular
            style="margin: 2rem"
            :rotate="180"
            :size="450"
            :width="25"
            :model-value="progressItems.value"
            color="blue"
          >
            {{ progressItems.value }} Qualifications
          </v-progress-circular>

          <v-progress-circular
            style="margin: 2rem"
            :rotate="180"
            :size="450"
            :width="25"
            :model-value="progressItems.value2"
            color="pink"
          >
            {{ progressItems.value2 }} Qualifications
          </v-progress-circular>
        </div>

        <div class="d-flex justify-center">
          <v-progress-linear
            v-model="knowledge"
            height="25"
            color="red"
            style="width: 90%"
          >
            <v-card-title>Progress Curriculum:</v-card-title>
            <strong style="color: red">{{ Math.ceil(knowledge) }}%</strong>
          </v-progress-linear>
        </div>
      </v-card>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, onMounted, reactive, ref } from "vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "DashboardDefault",
  components: { BaseLayout },
  setup() {
    const progressItems = reactive({
      interval: 0,
      value: 0,
      value2: 0,
    });
    const knowledge = ref(33);

    onBeforeMount(() => {
      clearInterval(progressItems.interval);
    });

    onMounted(() => {
      progressItems.interval = setInterval(() => {
        if (progressItems.value2 >= 100) {
          progressItems.value2 = 0;
          progressItems.value = 0;
          return;
        }
        progressItems.value += 10;
        progressItems.value2 += 25;
      }, 1000);
    });

    return { progressItems, knowledge };
  },
});
</script>
