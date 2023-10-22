// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import PerfogicCosmosCheckersPerfogicCosmoscheckersCosmoscheckers from './perfogic/cosmos-checkers/perfogic.cosmoscheckers.cosmoscheckers'
import PerfogicCosmosCheckersPerfogicCosmoscheckersLeaderboard from './perfogic/cosmos-checkers/perfogic.cosmoscheckers.leaderboard'


export default { 
  PerfogicCosmosCheckersPerfogicCosmoscheckersCosmoscheckers: load(PerfogicCosmosCheckersPerfogicCosmoscheckersCosmoscheckers, 'perfogic.cosmoscheckers.cosmoscheckers'),
  PerfogicCosmosCheckersPerfogicCosmoscheckersLeaderboard: load(PerfogicCosmosCheckersPerfogicCosmoscheckersLeaderboard, 'perfogic.cosmoscheckers.leaderboard'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
