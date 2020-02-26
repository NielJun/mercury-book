<template>
  <div class="home">
    <Tabs value="1" @on-click="changeTab">
      <TabPane  v-for="category of category_list" :key="category.id" :label="category.name" :id="category.id">
        <Card v-for="item of current_list" :key="item.question_id" :title="item.caption">
          <p>{{item.content}}</p>
          <p></p>
          <p>提问时间：{{item.create_time}} 提问者：{{item.author_name}}</p>
        </Card>
      </TabPane>
    </Tabs>
  </div>
</template>

<script>
export default {
  name: "home",
  data() {
    return {
      current_list: [],
      category_list:[]
    };
  },
   created() {
    this.fetchCategoryList()
  },
  methods: {
    fetchCategoryList() {
        let vm = this
        this.$http.get("/api/category/list").then(function(response){
              console.log(response);
              if (response.status != 200) {
                  vm.$Message.error("服务繁忙，请稍后重试");
                  return;
              }
              if (response.data.code === 0) {
                  console.log(response.data.data);
                  vm.category_list = response.data.data
                  if (vm.category_list.length > 0) {
                      vm.fetchQuestionList(vm.category_list[0].id)
                  }
              } else {
                  vm.$Message.error(response.data.message);
              }
            },
            function(response) {
                vm.$Message.error("服务繁忙，请稍后重试");
                console.log(response);
            }
        )
    },

    fetchQuestionList(id) {
      console.info(id);
      let vm = this
      this.$http.get("/api/question/list?category_id="+id).then(function(response){
              console.log(response);
              if (response.status != 200) {
                  vm.$Message.error("服务繁忙，请稍后重试");
                  return;
              }
              if (response.data.code === 0) {
                  console.log(response.data.data);
                  vm.current_list = response.data.data
              } else {
                  vm.$Message.error(response.data.message);
              }
            }, function(response) {
                vm.$Message.error("服务繁忙，请稍后重试");
                console.log(response);
            }
        )
    },

    changeTab:function (key) {
      const index = parseInt(key);
      this.fetchQuestionList(this.category_list[index].id)
    }
  }
};
</script>
