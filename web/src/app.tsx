import { Router, Route } from 'wouter-preact';
import { Home } from './pages';

export function App() {
  return (
    <Router>
      <Route path="/" component={Home} />
    </Router>
  )
}
