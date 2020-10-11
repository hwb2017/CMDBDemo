<template>
  <div :style="{ padding: '24px', background: '#fff', minHeight: '360px', margin: '24px 8px'}">
      <div class="content-header">
        <a-input :placeholder="searchItemPlaceholder" style="width: 320px " v-model="searchItemValue">
          <a-select slot="addonBefore" default-value="applicant" v-model="searchItemKey" style="width: 120px" @change="changePlaceholder">
            <a-select-option value="applicant">
              申请人
            </a-select-option>
            <a-select-option value="maintainer">
              维护人
            </a-select-option>
          </a-select>
        </a-input>
        <a-icon type="search" :style="{ fontSize: '28px', color: '#08c', margin: '2px 8px'}"/>
        <a-button type="primary" @click="openModal">添加</a-button>
    </div>  
    <a-card
      style="width:100%"
      :title="`${item.Applicant}的主机申请(${item.CreationTime})`"
      :tab-list="tabList"
      :active-tab-key="key"
      @tabChange="key => onTabChange(key, 'key')"
      v-for="item in vmLifecycles"
      :key="item._id"
    >
      <a-button-group slot="extra">
        <a-button type="primary" size="small" icon="edit"></a-button>
        <a-button type="danger" size="small" icon="delete"></a-button>
      </a-button-group>
      <div v-if="key == 'tab1'">
        <a-descriptions layout="vertical">
          <a-descriptions-item label="申请人">
            {{ item.Applicant }}
          </a-descriptions-item>
          <a-descriptions-item label="维护者">
            {{ item.Maintainer }}
          </a-descriptions-item>
          <a-descriptions-item label="申请到期时间">
            {{ item.VMLifecycleRules[0].actiontime }}
          </a-descriptions-item>         
        </a-descriptions>
      </div>
      <div v-else>
        关联主机信息...
      </div>
    </a-card>
    <a-modal v-model="visible" title="添加申请计划" @ok="handleOk">
      <a-form layout='vertical' :form="form">
        <a-form-item label='Title'>
          <a-input
            v-decorator="[
              'title',
              {
                rules: [{ required: true, message: 'Please input the title of collection!' }],
              }
            ]"
          />
        </a-form-item>
        <a-form-item label='Description'>
          <a-input
            type='textarea'
            v-decorator="['description']"
          />
        </a-form-item>
        <a-form-item class='collection-create-form_last-form-item'>
          <a-radio-group
            v-decorator="[
              'modifier',
              {
                initialValue: 'private',
              }
            ]"
          >
              <a-radio value='public'>Public</a-radio>
              <a-radio value='private'>Private</a-radio>
            </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>    
  </div>
</template>

<script>
import { mapActions } from 'vuex';
const placeholderMapping = {
  'applicant': '申请人',
  'maintainer': '维护人'
}
export default {
  data() {
    return {
      tabList: [
        {
          key: 'tab1',
          tab: '基本信息',
        },
        {
          key: 'tab2',
          tab: '关联主机',
        },
      ],
      key: 'tab1',
      searchItemPlaceholder: '请输入申请人',
      searchItemValue: '',
      searchItemKey: 'applicant',
      visible: false,
      form: {

      },
    };
  },
  computed: {
    vmLifecycles() {
      return this.$store.state.infra.vmLifecycles
    }
  },
  methods: {
    openModal() {
      this.visible = true;
    },
    handleOk(e) {
      console.log(e);
      this.visible = false;
    },
    onTabChange(key, type) {
      this[type] = key;
    },
    changePlaceholder(value) {
      this.searchItemPlaceholder = `请输入${placeholderMapping[value]}`;
    },
    ...mapActions({
      'getVMLifecycles': 'infra/getVMLifecycles'
    })
  },
  mounted: function() {
    this.getVMLifecycles();
  }
};
</script>

<style lang="less">
.content-header {
  padding: 0px 0px 24px;
  .ant-btn {
    float: right
  }
}
</style>