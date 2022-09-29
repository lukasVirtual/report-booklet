<template>
  <div v-if="$route.path == '/Berichtsheft'" class="text-center">
    <v-menu
      anchor="bottom"
      transition="slide-x-transition"
      @vnode-mounted="
        date = dateTextField;
        status = statusTextField;
      "
    >
      <template v-slot:activator="{ props }">
        <v-badge :content="stateTrue?.length ?? 0" color="accent">
          <v-btn
            @click="test"
            v-bind="props"
            icon
            style="margin-bottom: 30px"
            elevation="20"
            class="text-none"
          >
            <v-icon size="30">mdi-text-box-outline</v-icon>
          </v-btn>
        </v-badge>
      </template>
      <v-list width="400px" height="500">
        <v-list-item v-for="(quali, idx) in qualis" :key="idx">
          <template v-slot:prepend>
            <v-list-item-title>{{ quali.Title }}</v-list-item-title>
          </template>
          <template v-slot:append>
            <v-checkbox
              v-model="quali.State"
              @click="changeState(idx)"
              color="deep-purple accent-4"
            ></v-checkbox>
          </template>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { defineComponent, onMounted, ref } from "@vue/runtime-core";

export default defineComponent({
  name: "QualificationsDefault",
  props: {
    dateTextField: String,
    statusTextField: String,
  },
  setup() {
    const date = ref<string | undefined>("");
    const qualis = ref();
    const status = ref<string | undefined>("");
    const stateTrue = ref();
    const qualifications: any = [
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
      {
        title: "the first qualification",
        state: false,
      },
    ];

    onMounted(async () => {
      qualis.value = await dataService.getQualifications(date.value as string);
      stateTrue.value = qualis.value?.filter((v: any) => v.State == true);
    });

    const test = async () => {
      console.log(date.value);
      console.log(status.value);
      if (qualis.value?.length < 16 || qualis.value == undefined) {
        await dataService.insertQualifications(
          qualifications,
          date.value as string
        );
        console.log("empty");
      }
      qualis.value = await dataService.getQualifications(date.value as string);
    };

    const changeState = async (idx: number) => {
      qualis.value[idx].State = !qualis.value[idx].State;
      await dataService.insertQualifications(
        qualis.value,
        date.value as string
      );
      qualis.value = await dataService.getQualifications(date.value as string);
      stateTrue.value = qualis.value?.filter((v: any) => v.State == true);
    };

    return {
      test,
      changeState,
      status,
      date,
      stateTrue,
      qualifications,
      qualis,
    };
  },
});
</script>
