'use strict'

/**
 * @param {Function[]} functions
 * @return {Function}
 */
var compose = function(functions) {
    
    return function(x) {
        let result;

        if (functions.length == 0) {
            return x;
        }

        for (let i = functions.length - 1; i >= 0; i--) {
            const func = functions[i];
            if (result) {
                result = func(result);
            } else {
                result = func(x);
            }
        }

        return result;
    }
};

module.exports = compose;