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
        <a-button type="primary" @click="openModal('add')">添加</a-button>
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
        <a-button type="primary" size="small" icon="edit" @click="openModal('edit')"></a-button>
        <a-button type="danger" size="small" icon="delete"></a-button>
      </a-button-group>
      <div v-if="key == 'tab1'">
        <a-descriptions layout="vertical">
          <a-descriptions-item label="申请人">
            {{ item.applicant }}
          </a-descriptions-item>
          <a-descriptions-item label="维护者">
            {{ item.maintainer }}
          </a-descriptions-item>
          <a-descriptions-item label="申请到期时间">
            {{ item.vmlifecyclerules[0].actiontime }}
          </a-descriptions-item>         
        </a-descriptions>
      </div>
      <div v-else>
        关联主机信息...
      </div>
    </a-card>
    <a-modal v-model="visible" :title="modalTitle" @ok="handleOk">
      <a-form layout='vertical' :form="form">
        <a-form-item label='申请人'>
          <a-input
            v-decorator="[
              'applicant',
              {
                rules: [{ required: true, message: '请输入申请人名称' }],
              }
            ]"
          />
        </a-form-item>
        <a-form-item label='维护人'>
          <a-input
            v-decorator="['maintainer']"
          />
        </a-form-item>
        <a-form-item 
          v-for="(k,index) in form.getFieldValue('vmLifecycleOps')"
          :key="k"
          :required=false
          :label="index === 0 ? '生命周期策略' : ''"
        >
          <a-select default-value="stop" style="width: 100px">
            <a-select-option value="stop">停机</a-select-option>
            <a-select-option value="destroy">销毁</a-select-option>
          </a-select>
          <a-date-picker
            v-decorator="[
              `date-time-picker[${k}]`, 
              {
                rules: [{ type: 'object', required: true, message: '请选择时间' }],
              }
            ]"
            show-time
            format="YYYY-MM-DD HH:mm:ss"
            :style="{ margin: '0 8px' }"
          />
          <a-icon
             type="plus-circle"
             @click="() => add()"
             :style="{ margin: '0 8px' }"
          />
          <a-icon
             v-if="form.getFieldValue('vmLifecycleOps').length > 1"
             type="minus-circle"
             @disabled="form.getFieldValue('vmLifecycleOps').length === 1"
             @click="() => remove(k)"
             :style="{ margin: '0 8px' }"
          />
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
let id = 0;
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
      modalTitle: '',
    };
  },
  beforeCreate() {
    this.form = this.$form.createForm(this, {name: 'application_plan'});
    this.form.getFieldDecorator('vmLifecycleOps', { initialValue: [0], preserve: true });
  },
  computed: {
    vmLifecycles() {
      return this.$store.state.infra.vmLifecycles
    }
  },
  methods: {
    add() {
      const { form } = this;
      const ops = form.getFieldValue('vmLifecycleOps');
      const nextOps = ops.concat(++id);
      form.setFieldsValue({
        vmLifecycleOps: nextOps,
      });
    },
    remove(k) {
      const { form } = this;
      const ops = form.getFieldValue('vmLifecycleOps');
      if (ops.length === 1) {
        return;
      }
      form.setFieldsValue({
        vmLifecycleOps: ops.filter(key => key !== k),
      });
    },
    openModal(action) {
      this.visible = true;
      if (action=="add") {
          this.modalTitle = "添加申请计划"
      } else if (action=="edit") {
          this.modalTitle = "编辑申请计划"
      }
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