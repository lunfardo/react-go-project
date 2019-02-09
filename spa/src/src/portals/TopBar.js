import React, { Component } from 'react'
import ReactDOM from 'react-dom';


const topbar = document.getElementById('top-bar');

class TopBar extends Component {
  constructor(props) {
    super(props);
    this.el = document.createElement('div');
  }

  componentDidMount() {
    topbar.appendChild(this.el);
  }

  componentWillUnmount() {
    topbar.removeChild(this.el);
  }

  render() {
    return ReactDOM.createPortal(
      this.props.children,
      this.el,
    );
  }
}

export default TopBar