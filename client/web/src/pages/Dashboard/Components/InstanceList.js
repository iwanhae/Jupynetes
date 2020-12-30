import React, { Component } from 'react'
import { Button, Card, Row, Col} from 'antd';
import Instance from './Instance.js';

const style = { background: '#0092ff', padding: '8px 0' };

export default class InstanceList extends Component {
    render() {
        return (
            <div>
                <Row gutter={16}>
                    <Col className="gutter-row" span={6}>
                    </Col>
                    <Col className="gutter-row" span={6}>
                    </Col>
                    <Col className="gutter-row" span={6}>
                    </Col>
                    <Col className="gutter-row" span={6}>
                        <div style={style}>
                            <Button danger>인스턴스 생성</Button>
                        </div>
                    </Col>
                </Row>
                <Card title="Card Title">
                    <Instance/>
                    <Instance/>
                    <Instance/>
                    <Instance/>
                    <Instance/>
                </Card>
            </div>
        )
    }
}
