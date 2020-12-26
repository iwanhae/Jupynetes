import React from "react";
import GridItem from "components/Grid/GridItem.js";
import GridContainer from "components/Grid/GridContainer.js";

import Button from "components/CustomButtons/Button.js";

export default function Instane(props) {

  const statusStyle={
    height: "50px",
    width: "50px",
    backgroundColor: "#F00",
    borderRadius: "50%",
    display: "inline-block",
  }

  const deleteButtonStyle={
    height: "50px",
    width: "50px",
    backgroundColor: "#992222",
  }

  const accessButtonStyle={
    height: "50px",
    width: "50px",
    backgroundColor: "#009955",
  }

    return<GridContainer>
            <GridItem xs={12} sm={6} md={5}>
                <span class="dot" style={statusStyle}></span>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
                <Button type="button" style={deleteButtonStyle}>
                    삭제
                </Button>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
                <Button type="button" style={accessButtonStyle}>
                    접속
                </Button>
            </GridItem>
        </GridContainer>
}