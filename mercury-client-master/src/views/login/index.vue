<template>
  <div class="login">
    <div class="login-box">
      <Card title="登录">
        <Form>
          <FormItem prop="username">
            <Input type="text" v-model="form.user" placeholder="Username">
              <Icon type="ios-person-outline" slot="prepend"></Icon>
            </Input>
          </FormItem>
          <FormItem prop="password">
            <Input type="password" v-model="form.password" placeholder="Password">
              <Icon type="ios-lock-outline" slot="prepend"></Icon>
            </Input>
          </FormItem>
          <Button type="primary" long @click="submit">登录</Button>
          <router-link
            class="ivu-btn ivu-btn-default ivu-btn-long"
            style="margin-top:10px;"
            to="/register"
          >注册</router-link>
          <div style="margin-top:10px;text-align:right;">
            <router-link class="ivu-btn ivu-btn-text" to="/">返回首页</router-link>
          </div>
        </Form>
      </Card>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      form: {
        user: "",
        password: ""
      }
    };
  },
  methods: {
    async submit() {
        let res = await this.$http.post("/api/user/login", this.form);
        if (res.status!=200){
            this.$Message.error("登陆失败 网络错误");
            return
        }
        if (res.data.code === 0) {
            this.$Message.success("登陆成功");
            this.$router.push("/");
        }else{
            this.$Message.error(res.data.message);
        }
    }
  }
};
</script>

<style lang="scss" scoped src="./style.scss">
</style>
