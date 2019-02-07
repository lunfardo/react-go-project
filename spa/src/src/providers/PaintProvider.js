import React, { Component } from 'react';
import { PaintContext } from '../AppContexts'


class PaintProvider extends Component {
    state = {
      paintsList: []
    }

    async loadPaintsFromAPI(){
      return fetch(`//${window.location.host}/api/v1/paints`)
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
        this.loadPaintsFromAPI().then((result)=>{
            this.setState({paintsList: result})    
        })          
    }    
    render() {
      return (
        <PaintContext.Provider value={this.state}>
          {this.props.children}
        </PaintContext.Provider>
      )}
}

export default PaintProvider