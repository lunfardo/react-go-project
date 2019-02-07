import React, { Component } from 'react';
import { CommentContext } from '../AppContexts'


class CommentProviders extends Component {
    state = {
      somestate : "okayish"
    }
    render() {
      return (
        <PaintersContext.Provider value={this.state}>
          {this.props.children}
        </PaintersContext.Provider>
      )}
}

export default CommentProviders