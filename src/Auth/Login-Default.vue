<template>
  <v-app>
    <v-main
      v-if="$route.path == '/Login'"
      class="text-center"
      style="display: flex; align-items: center; margin: auto"
    >
      <v-form
        tile
        ref="valid"
        style="
          border: 2px solid black;
          height: 400px;
          width: 700px;
          display: flex;
          justify-content: center;
          align-items: center;
          border-radius: 20px;
        "
      >
        <v-row justify="center">
          <v-card-title>Login</v-card-title>
          <v-col cols="12">
            <v-text-field
              prepend-icon="mdi-account"
              label="Username"
              type="text"
              color="cyan"
              style="padding-left: 25%; padding-right: 25%"
              :rules="[rules.required]"
              variant="underlined"
              v-model="validation.username"
            ></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field
              prepend-icon="mdi-lock"
              label="Password"
              type="password"
              style="padding-left: 25%; padding-right: 25%"
              variant="underlined"
              color="cyan"
              :rules="[rules.required, rules.minLen]"
              v-model="validation.password"
              @keyup.enter="login"
            ></v-text-field>
          </v-col>

          <v-divider length="700px"></v-divider>
          <v-col>
            <div class="text-right">
              <v-card-actions>
                <v-btn @click="login">Login</v-btn>

                <v-btn color="red" rounded @click="reset">reset</v-btn>
              </v-card-actions>
            </div>
          </v-col>
        </v-row>
      </v-form>
      <v-overlay :model-value="loading" class="align-center justify-center">
        <v-progress-circular indeterminate size="64"></v-progress-circular>
      </v-overlay>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { loginService } from "@/handler/loginHandler";
import { defineComponent, reactive, ref } from "@vue/runtime-core";
import { io } from "socket.io-client";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";

export default defineComponent({
  name: "LoginDefault",

  setup() {
    const rules = {
      required: (value: string) => !!value || "Required",
      minLen: (v: string) => v.length >= 8 || "Min. 8 Characters",
    };
    const validation = reactive({
      username: "",
      password: "",
    });
    const socket = io("http://localhost:7000");
    const router = useRouter();
    const toastMessage = useToast();

    const loading = ref(false);

    const login = async () => {
      loading.value = true;
      let registerRequest = ref(false);

      if (validation.username !== "" && validation.password.length >= 8) {
        try {
          registerRequest.value = await loginService.login(
            validation.username,
            validation.password
          );
          if (registerRequest.value) {
            loading.value = false;
            socket.emit("login", {
              user_name: await loginService.getUser(),
              user_id: await loginService.getUserID(),
            });
            for (const route of router.getRoutes()) {
              if (route.meta.requiresAuth) {
                let err = await router.push({
                  name: route.name,
                });

                if (err == undefined) return;
              }
              router.push({ path: "/Login" });
            }
          } else {
            toastMessage.error(
              "something went wrong. Check your password and username."
            );
            loading.value = false;
          }
        } catch (e) {
          toastMessage.error(
            "something went wrong. Check your password and username."
          );
          loading.value = false;
        }
        console.log(registerRequest.value);
      } else {
        toastMessage.error("Something went wrong. Wrong username or password");
        loading.value = false;
      }
      loading.value = false;
    };

    const valid = ref(null);
    const reset = () => {
      const obj = valid.value as any;
      if (obj !== undefined) obj.reset();
    };
    return { rules, validation, valid, loading, login, reset };
  },
});
</script>
