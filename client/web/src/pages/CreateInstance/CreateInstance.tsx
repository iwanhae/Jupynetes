import React from 'react';
import 'antd/dist/antd.css';
import '../../index.css';
import { Button, Modal, Form, Input, Dropdown, Menu } from 'antd';
import CSS from 'csstype';
import { DownOutlined, UserOutlined } from '@ant-design/icons';
import appState from '../../AppState';
import { ServerObject } from '../../model/ServerObject';

const mypageStyle: CSS.Properties = {
  margin: '1rem',
  float: 'right',
};

const layout = {
  labelCol: {
    span: 6,
  },
  wrapperCol: {
    span: 12,
  },
};
const tailLayout = {
  wrapperCol: {
    offset: 6,
    span: 12,
  },
};

const menu = (
  <Menu onClick={handleMenuClick}>
    <Menu.Item key="1" icon={<UserOutlined />}>
      상
    </Menu.Item>
    <Menu.Item key="2" icon={<UserOutlined />}>
      중
    </Menu.Item>
    <Menu.Item key="3" icon={<UserOutlined />}>
      하
    </Menu.Item>
  </Menu>
);

const storageMenu = (
  <Menu onClick={handleMenuClick}>
    <Menu.Item key="1" icon={<UserOutlined />}>
      상
    </Menu.Item>
    <Menu.Item key="2" icon={<UserOutlined />}>
      중
    </Menu.Item>
    <Menu.Item key="3" icon={<UserOutlined />}>
      하
    </Menu.Item>
  </Menu>
);

function handleMenuClick() {}

const onFinish = (values: String) => {
  console.log('Success:', values);
};

const onFinishFailed = () => {};

const CreateInstance = () => {
  const [visible, setVisible] = React.useState(false);
  const [confirmLoading, setConfirmLoading] = React.useState(false);
  const [modalText, setModalText] = React.useState('Content of the modal');

  const showModal = () => {
    setVisible(true);
  };

  const handleOk = () => {
    appState.addServer(new ServerObject());
    setModalText('The modal will be closed after two seconds');
    setConfirmLoading(true);
    setTimeout(() => {
      setVisible(false);
      setConfirmLoading(false);
    }, 2000);
  };

  const handleCancel = () => {
    console.log('Clicked cancel button');
    setVisible(false);
  };

  return (
    <>
      <Button danger style={mypageStyle} onClick={showModal}>
        인스턴스생성
      </Button>
      <Modal
        title="인스턴스 생성"
        visible={visible}
        onOk={handleOk}
        confirmLoading={confirmLoading}
        onCancel={handleCancel}
      >
        <Form
          {...layout}
          name="basic"
          initialValues={{
            remember: true,
          }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
        >
          <Form.Item
            label="서버이름"
            name="servername"
            rules={[
              {
                required: true,
                message: '서버명을 입력해 주세요.',
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="서버설명"
            name="description"
            rules={[
              {
                required: true,
                message: '간단한 설명을 입력해 주세요.',
              },
            ]}
          >
            <Input.Password />
          </Form.Item>
          <Form.Item
            label="Flavor"
            name="flavor"
            rules={[
              {
                required: true,
                message: '서버 사양을 선택해 주세요.',
              },
            ]}
          >
            <Dropdown overlay={menu}>
              <Button>
                선택 <DownOutlined />
              </Button>
            </Dropdown>
          </Form.Item>
          <Form.Item
            label="Storage"
            name="storage"
            rules={[
              {
                required: true,
                message: '디스크 사양을 선택해 주세요.',
              },
            ]}
          >
            <Dropdown overlay={storageMenu}>
              <Button>
                선택 <DownOutlined />
              </Button>
            </Dropdown>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default CreateInstance;
