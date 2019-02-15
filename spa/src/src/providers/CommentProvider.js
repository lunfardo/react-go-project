import React, { Component } from 'react';
import { CommentContext } from '../AppContexts'


class CommentProvider extends Component {
    state = {
      somestate : "okayish",
      alert: this.alert,
    }

    alert(){
      alert('asdfjdk;')
    }
    render() {
      return (
        <CommentContext.Provider value={this.state}>
          {this.props.children}
        </CommentContext.Provider>
      )}
}

export default CommentProvider