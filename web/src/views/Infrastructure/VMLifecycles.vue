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
    <a-table
      :columns="columns"
      :data-source="vmLifecycles"
      :row-key="record => record._id"
      :pagination=false
    >  
      <div slot="ApplicationPlanTitle" slot-scope="record">
        {{ record.applicant }}的主机申请({{ record.createtime | datefmt }})
      </div>
      <div slot="Deadline" slot-scope="record">
        {{ record.vmlifecyclerules[0].actiontime | datefmt }}
      </div>
      <a-button-group slot="Operation" slot-scope="record">
        <a-button type="primary" size="small" icon="edit" @click="openModal('edit', record)"></a-button>
        <a-button type="danger" size="small" icon="delete" @click="deleteApplicationPlan(record._id)"></a-button>
      </a-button-group>
    </a-table>
    <a-modal v-model="visible" :title="modalTitle" @ok="handleOk">
      <a-form layout='vertical' :form="form">
        <a-form-item label='申请人'>
          <a-input v-decorator="[
            'applicant',
            {
              rules: [{ required: true, message: '请输入申请人名称' }],
            }]"
          />
        </a-form-item>
        <a-form-item label='维护人'>
          <a-input
            v-decorator="['maintainer']"
          />
        </a-form-item>
        <a-form-item 
          v-for="(k,index) in form.getFieldValue('vmLifecycleRuleKeys')"
          :key="k"
          :required=false
          :label="index === 0 ? '生命周期策略' : ''"
        >
          <a-select 
            style="width: 100px"
            v-decorator="[`operation[${k}]`]"
          >
            <a-select-option value="stop">停机</a-select-option>
            <a-select-option value="destroy">销毁</a-select-option>
          </a-select>
          <a-date-picker
            v-decorator="[
              `actionTime[${k}]`, 
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
             v-if="form.getFieldValue('vmLifecycleRuleKeys').length > 1"
             type="minus-circle"
             @disabled="form.getFieldValue('vmLifecycleRuleKeys').length === 1"
             @click="() => remove(k)"
             :style="{ margin: '0 8px' }"
          />
        </a-form-item>  
        <a-form-item label="关联主机">
          <a-select
            mode="multiple"
            style="width: 100%"
            placeholder="请输入主机ID精确搜索"
            v-decorator="['vm_ids', { initialValue: [] }]"
            option-label-prop="label"
            :filter-option="false"
            :not-found-content="fetching ? undefined : null"
            @search="searchVMByID"
          >
            <a-spin v-if="fetching" slot="notFoundContent" size="small" />
            <a-select-option v-for="vm in filteredVMs" :key="vm._id" :label="vm._id">
              <p><strong>实例ID: </strong>{{ vm._id }}</p>
              <p><strong>实例名称: </strong>{{ vm.instance_name }}</p>
            </a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>    
  </div>
</template>

<script>
import { mapActions } from 'vuex';
import moment from "moment";
import debounce from "lodash/debounce";
const placeholderMapping = {
  'applicant': '申请人',
  'maintainer': '维护人'
}
const columns = [
  {
    key: 'ApplicationPlanTitle',
    title: '申请计划名称',
    scopedSlots: { customRender: 'ApplicationPlanTitle' },
  },
  {
    key: 'Applicant',
    dataIndex: 'applicant',
    title: '申请人',
  },
  {
    key: 'Maintainer',
    dataIndex: 'maintainer',
    title: '维护人',
  },
  {
    key: 'Deadline',
    title: '申请到期时间',
    scopedSlots: { customRender: 'Deadline' },
  },
  {
    key: 'Operation',
    title: '操作',
    scopedSlots: { customRender: 'Operation' },
  }
]
export default {
  data() {
    this.searchVMByID = debounce(this.searchVMByID);
    this.lastFetchId = 0;
    return {
      searchItemPlaceholder: '请输入申请人',
      searchItemValue: '',
      searchItemKey: 'applicant',
      visible: false,
      modalTitle: '',
      modalType: '',
      columns,
      filteredVMs: [],
      fetching: false
    };
  },
  beforeCreate() {
    this.form = this.$form.createForm(this, {name: 'application_plan'});
    this.form.getFieldDecorator('vmLifecycleRuleKeys', { initialValue: [0], preserve: true });
    this.form.getFieldDecorator('id', { preserve: true });
  },
  watch: {
    vmLifecycles() {
       this.getVMLifecycles();
    }
  },
  computed: {
    vmLifecycles() {
      return this.$store.state.infra.vmLifecycles
    },
  },
  methods: {
    add() {
      const { form } = this;
      const keys = form.getFieldValue('vmLifecycleRuleKeys');
      const nextID = Math.max(...form.getFieldValue('vmLifecycleRuleKeys'))+1;
      const nextKeys = keys.concat(nextID);
      form.setFieldsValue({
        vmLifecycleRuleKeys: nextKeys,
      });
    },
    remove(k) {
      const { form } = this;
      const keys = form.getFieldValue('vmLifecycleRuleKeys');
      if (keys.length === 1) {
        return;
      }
      form.setFieldsValue({
        vmLifecycleRuleKeys: keys.filter(key => key !== k),
      });
    },
    openModal(action, currentRow) {
      this.visible = true;
      this.filteredVMs = [];
      const { form } = this; 
      if (action=="add") {
          this.modalTitle = "添加申请计划"
          this.modalType = "add"
          form.resetFields();    
      } else if (action=="edit") {
          this.modalTitle = "编辑申请计划"
          this.modalType = "edit"
          this.$nextTick(() => {
            form.setFieldsValue({ 
              id: currentRow._id,
              applicant: currentRow.applicant,
              maintainer: currentRow.maintainer,
              vm_ids: currentRow.vmids
            });
            var vmLifecycleRuleItems = {};
            var vmLifecycleRuleKeys = [];
            for (let i = 0; i < currentRow.vmlifecyclerules.length; i++) {
              form.getFieldDecorator(`operation[${i}]`);
              form.getFieldDecorator(`actionTime[${i}]`);
              vmLifecycleRuleItems[`operation[${i}]`] = currentRow.vmlifecyclerules[i].operation;
              vmLifecycleRuleItems[`actionTime[${i}]`] = moment.unix(currentRow.vmlifecyclerules[i].actiontime);
              vmLifecycleRuleKeys.push(i);
            }
            form.setFieldsValue({ vmLifecycleRuleKeys: vmLifecycleRuleKeys });
            form.setFieldsValue(vmLifecycleRuleItems);
          });
          for (let vmid of currentRow.vmids) {
            this.getVirtualMachine(vmid)
            .then(body => {
              this.filteredVMs.push({
                _id: body.data.data[0]["_id"],
                instance_name: body.data.data[0]["instance_name"]
              })
            })
          };
      }
    },
    handleOk() {
      const { form } = this;
      const vmLifecycleRules = [];
      for ( let i of form.getFieldValue('vmLifecycleRuleKeys')) {
          const vmLifecycleRule = {
            "operation": form.getFieldValue(`operation[${i}]`),
            "action_time": form.getFieldValue(`actionTime[${i}]`).unix()
          };
          vmLifecycleRules.push(vmLifecycleRule);
      }
      const payload = {
        "maintainer": form.getFieldValue('maintainer'),
        "applicant": form.getFieldValue('applicant'),
        "vm_lifecycle_rules": vmLifecycleRules,
        "vm_ids": form.getFieldValue('vm_ids'),
      };
      const createPayload = payload, updatePayload = payload;
      updatePayload["id"] = form.getFieldValue('id');
      if (this.modalType == "add") {
          this.createApplicationPlan(createPayload);
      } else if (this.modalType == "edit") {
          this.updateApplicationPlan(updatePayload);
      }
      this.visible = false;
      form.resetFields();
      this.filteredVMs = [];
    },
    createApplicationPlan(payload) {
      this.createVMLifecycles(payload);
    },
    deleteApplicationPlan(id) {
      this.deleteVMLifecycles(id);
      this.getVMLifecycles();
    },  
    updateApplicationPlan(payload) {
      this.updateVMLifecycles(payload);
    },
    changePlaceholder(value) {
      this.searchItemPlaceholder = `请输入${placeholderMapping[value]}`;
    },
    searchVMByID(value) {
      this.lastFetchId += 1;
      const fetchId = this.lastFetchId;
      this.fetching = true;
      this.getVirtualMachine(value)
        .then(body => {
          if (fetchId !== this.lastFetchId) {
            return;
          }
          this.filteredVMs.push({
            _id: body.data.data[0]["_id"],
            instance_name: body.data.data[0]["instance_name"]
          })
          this.fetching = false;
        });
    },
    ...mapActions({
      'getVMLifecycles': 'infra/getVMLifecycles',
      'createVMLifecycles': 'infra/createVMLifecycles',
      'updateVMLifecycles': 'infra/updateVMLifecycles',
      'deleteVMLifecycles': 'infra/deleteVMLifecycles',
      'getVirtualMachine': 'infra/getVirtualMachine'
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