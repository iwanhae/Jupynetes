import React from "react";
// react plugin for creating charts
// @material-ui/core
import { makeStyles } from "@material-ui/core/styles";
// @material-ui/icons
// core components
import GridItem from "components/Grid/GridItem.js";
import GridContainer from "components/Grid/GridContainer.js";
import Instance from "./components/Instance.js"
import FloatingButtons from './components/FloatingButtons';

const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      position: "fixed",
    },  fab: {
      position: "fixed",
      top: "300px",
      left: "300px"
    },
  },
}));

export default function Dashboard() {
  const classes = useStyles();
  return (
    <div>
      <FloatingButtons style={classes.fab}/>
      <GridContainer>
        <GridItem xs={12} sm={6} md={3}>
          <Instance>사과서버</Instance>
        </GridItem>
        <GridItem xs={12} sm={6} md={3}>
          <Instance>감귤서버</Instance>
        </GridItem>
        <GridItem xs={12} sm={6} md={3}>
          <Instance>딸기서버</Instance>
        </GridItem>
        <GridItem xs={12} sm={6} md={3}>
          <Instance>멜론서버</Instance>
        </GridItem>
        <GridItem xs={12} sm={6} md={3}>
          <Instance>감귤서버</Instance>
        </GridItem>
      </GridContainer>
    </div>
  );
}

