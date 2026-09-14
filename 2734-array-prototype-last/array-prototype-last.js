/**
 * @return {null|boolean|number|string|Array|Object}
 */
Array.prototype.last = function() {
    let l = this.length;
    if (l == 0) {
        return -1;
    }

    return this[l-1]
};