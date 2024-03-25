Module[ 'onRuntimeInitialized' ] = onRuntimeInitialized;

const add = Module.cwrap( 'add', 'number', [ 'number', 'number' ] );

function onRuntimeInitialized () {
    log( add( 40, 2 ) );
}
