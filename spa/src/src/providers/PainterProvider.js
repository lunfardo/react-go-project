import React, { Component } from 'react';
import { PainterContext } from '../AppContexts'


class PainterProvider extends Component {
    state = {
      paintersList: []
    }

    async loadPaintersFromAPI(){
      return fetch(`//${window.location.host}/api/v1/painters`)
      .then((response) => {
        if(response.ok) {
          return response.json()
        } else {
          return 'not ok'
          }
      })
      .catch(function(error) {
        this.setState({apimessage: 'Fetch failed:' + error.message})
      });  
    }

    componentDidMount() {
        this.loadPaintersFromAPI().then((result)=>{
            this.setState({paintersList: result})    
        })          
    }    
    render() {
      return (
        <PainterContext.Provider value={this.state}>
          {this.props.children}
        </PainterContext.Provider>
      )}
}

export default PainterProvider