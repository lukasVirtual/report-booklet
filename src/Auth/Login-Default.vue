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
          border: 1px solid black;
          height: 500px;
          width: 800px;
          display: flex;
          justify-content: center;
          align-items: center;
        "
      >
        <v-row justify="center">
          <v-card-title>Login</v-card-title>
          <v-col cols="12">
            <v-text-field
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

          <v-divider length="800px"></v-divider>
          <v-col>
            <div class="text-right">
              <v-card-actions>
                <v-btn rounded @click="login">Login</v-btn>

                <v-btn color="red" rounded @click="reset">reset</v-btn>
              </v-card-actions>
            </div>
          </v-col>
        </v-row>
      </v-form>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { loginService } from "@/handler/loginHandler";
import { defineComponent, reactive, ref } from "vue";
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

    const router = useRouter();
    const toastMessage = useToast();

    const login = async () => {
      let registerRequest = ref(false);

      if (validation.username !== "" && validation.password.length >= 8) {
        try {
          registerRequest.value = await loginService.login(
            validation.username,
            validation.password
          );
          if (registerRequest.value)
            for (const route of router.getRoutes()) {
              if (route.meta.requiresAuth) {
                let err = await router.push({
                  name: route.name,
                });

                if (err == undefined) return;
              }
              router.push({ path: "/Login" });
            }
          else
            toastMessage.error(
              "something went wrong. Check your password and username."
            );
        } catch (e) {
          toastMessage.error(
            "something went wrong. Check your password and username."
          );
        }
        console.log(registerRequest.value);
      } else
        toastMessage.error("Something went wrong. Wrong username or password");
    };

    const valid = ref(null);
    const reset = () => {
      const obj = valid.value as any;
      if (obj !== undefined) obj.reset();
    };
    return { rules, validation, valid, login, reset };
  },
});
</script>
