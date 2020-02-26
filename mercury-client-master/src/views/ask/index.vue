<template>
  <div>
    <Card title="提问">
      <Form label-position="top" :model="form">
        <FormItem prop="title" label="标题">
          <Input type="text" v-model="form.caption" placeholder="请输入标题"/>
        </FormItem>
        <FormItem prop="category" label="类别">
          <Select v-model="form.category_id" placeholder="请选择类别">
            <Option v-for="category of category_list" :value="category.id">{{category.name}}</Option>
          </Select>
        </FormItem>
        <FormItem prop="centent" label="内容">
          <Input type="textarea" v-model="form.content" placeholder="请输入内容"/>
        </FormItem>
        <Button type="primary" @click="submit">提交</Button>
      </Form>
    </Card>
  </div>
</template>

<script>
    export default {
        name: "Ask",
        data() {
            return {
                form: {
                    caption: "",
                    category_id: 1,
                    content: "",
                }, category_list: []
            };
        }, created() {
            this.fetchCategoryList()
        },
        methods: {
            async submit() {
                let res = await this.$http.post("/api/ask/submit", this.form);
                console.log(res);
                if (res.status != 200) {
                    this.$Message.error("提问失败，网络错误");
                    return;
                }
                if (res.data.code === 0) {
                    this.$Message.success("提问成功");
                    this.$router.push("/");
                } else {
                    // 这个1008错误码最好是使用常量定义
                    this.$Message.error(res.data.message);
                    if (res.data.code === 1008) {
                        this.$router.push("/login")
                    }
                }
            },
            fetchCategoryList() {
                this.category_list = [{id: 1, name: "技术"}];
                let vm = this;
                this.$http.get("api/category/list").then(function (response) {
                    console.log(response);
                    if (response.status != 200) {
                        vm.$Message.error("获取列表错误，网络有问题");
                        return;
                    }
                    if (response.data.code === 0) {
                        vm.category_list = response.data.data;
                    } else {
                        vm.$Message.error(response.data.message);
                    }
                });
            }
        }
    };
</script>

<style lang="scss" scoped>
</style>
