// Variables globales para la máquina de Turing
let tape = [];
let headPosition = 0;
let currentState = '';
let transitionRules = new Map(); // Mapa: 'estado,simbolo' -> {nuevoSimbolo, movimiento, nuevoEstado}
let stepsCount = 0;
let isRunning = false; // Para controlar la ejecución automática

// Elementos del DOM
const inputTapeElement = document.getElementById('inputTape');
const rulesFileElement = document.getElementById('rulesFile');
const loadRulesButton = document.getElementById('loadRules');
const resetMachineButton = document.getElementById('resetMachine');
const stepMachineButton = document.getElementById('stepMachine');
const runMachineButton = document.getElementById('runMachine');
const tapeDisplay = document.getElementById('tape');
const headDisplay = document.getElementById('head');
const currentStateDisplay = document.getElementById('currentState');
const stepsCountDisplay = document.getElementById('stepsCount');
const executionLog = document.getElementById('executionLog');

// Símbolo de blanco
const BLANK_SYMBOL = '_';

// Función para cargar las reglas de transición desde un archivo
async function loadRules() {
    const file = rulesFileElement.files[0];
    if (!file) {
        alert('Por favor, selecciona un archivo de reglas (output.txt).');
        return;
    }

    const reader = new FileReader();
    reader.onload = (e) => {
        const content = e.target.result;
        parseRules(content);
        alert('Reglas cargadas exitosamente.');
        initializeMachine(); // Inicializar la máquina después de cargar las reglas
    };
    reader.readAsText(file);
}

// Función para parsear las reglas de transición
function parseRules(content) {
    transitionRules.clear();
    const lines = content.split('\n').filter(line => line.trim() !== '');
    lines.forEach(line => {
        // Formato esperado: estado_actual simbolo_leido simbolo_escrito movimiento nuevo_estado
        const parts = line.trim().split(' ');
        if (parts.length === 5) {
            const [q, s, newS, move, nextQ] = parts;
            transitionRules.set(`${q},${s}`, { newS, move, nextQ });
        } else {
            console.warn(`Línea de regla inválida: ${line}`);
        }
    });
    console.log('Reglas de Transición:', transitionRules);
}

// Función para inicializar la máquina de Turing
function initializeMachine() {
    const initialTapeString = inputTapeElement.value.trim();
    tape = initialTapeString.split('');
    headPosition = 0;
    currentState = 'a'; // Estado inicial por defecto
    stepsCount = 0;
    executionLog.textContent = ''; // Limpiar log
    renderTape();
    updateMachineDisplay();
    logExecution(`Máquina inicializada. Cinta: ${tape.join('')}, Estado: ${currentState}, Cabezal: ${headPosition}`);
    enableControls(true);
}

// Función para renderizar la cinta en el DOM
function renderTape() {
    tapeDisplay.innerHTML = '';
    // Asegurarse de que haya suficiente cinta a la izquierda y derecha
    const displayTape = [...tape];
    const padding = 10; // Celdas de blanco a cada lado para visualización
    for (let i = 0; i < padding; i++) {
        displayTape.unshift(BLANK_SYMBOL);
        displayTape.push(BLANK_SYMBOL);
    }

    const startIndex = headPosition + padding - Math.floor(tapeDisplay.offsetWidth / 80); // Ajuste para centrar el cabezal
    const endIndex = headPosition + padding + Math.ceil(tapeDisplay.offsetWidth / 80);

    for (let i = 0; i < displayTape.length; i++) {
        const cell = document.createElement('div');
        cell.classList.add('tape-cell');
        cell.textContent = displayTape[i];
        if (i === headPosition + padding) { // La celda actual del cabezal
            cell.classList.add('active');
        }
        tapeDisplay.appendChild(cell);
    }

    // Ajustar el scroll para centrar el cabezal
    const activeCell = tapeDisplay.querySelector('.tape-cell.active');
    if (activeCell) {
        tapeDisplay.scrollLeft = activeCell.offsetLeft - (tapeDisplay.offsetWidth / 2) + (activeCell.offsetWidth / 2);
        
        // Mover el cabezal visualmente
        // Calcular la posición left del cabezal para que esté sobre la celda activa
        // El cabezal debe estar centrado sobre la celda activa
        const headOffset = activeCell.offsetLeft + (activeCell.offsetWidth / 2) - (headDisplay.offsetWidth / 2);
        headDisplay.style.left = `${headOffset}px`;
    }
}

// Función para actualizar el estado de la máquina en el DOM
function updateMachineDisplay() {
    currentStateDisplay.textContent = currentState;
    stepsCountDisplay.textContent = stepsCount;
    // La posición del cabezal visual se maneja en renderTape
}

// Función para registrar la ejecución
function logExecution(message) {
    executionLog.textContent += message + '\n';
    executionLog.scrollTop = executionLog.scrollHeight; // Scroll al final
}

// Función para ejecutar un paso de la máquina de Turing
function step() {
    if (currentState === 'z') { // 'z' es el estado de parada por defecto
        logExecution('Máquina detenida: Estado final "z" alcanzado.');
        isRunning = false;
        enableControls(true);
        return;
    }

    // Asegurarse de que la cinta tenga un símbolo en la posición del cabezal
    while (headPosition >= tape.length) {
        tape.push(BLANK_SYMBOL);
    }
    while (headPosition < 0) {
        tape.unshift(BLANK_SYMBOL);
        headPosition = 0; // Ajustar la posición del cabezal después de unshift
    }

    const currentSymbol = tape[headPosition];
    const ruleKey = `${currentState},${currentSymbol}`;
    const rule = transitionRules.get(ruleKey);

    if (!rule) {
        logExecution(`Error: No se encontró una regla para el estado ${currentState} y símbolo ${currentSymbol}.`);
        logExecution('Máquina detenida debido a un estado no definido.');
        isRunning = false;
        enableControls(true);
        return;
    }

    const { newS, move, nextQ } = rule;

    // Aplicar la regla
    const oldTape = [...tape]; // Copia para el log
    const oldHeadPos = headPosition;
    const oldState = currentState;

    tape[headPosition] = newS;
    if (move === 'd') {
        headPosition++;
    } else if (move === 'i') {
        headPosition--;
    }
    currentState = nextQ;
    stepsCount++;

    logExecution(`Paso ${stepsCount}: Estado: ${oldState}, Símbolo: ${currentSymbol} -> Nueva Cinta: ${tape.join('')}, Nuevo Estado: ${currentState}, Nueva Posición: ${headPosition}`);

    renderTape();
    updateMachineDisplay();

    if (currentState === 'z' && isRunning) {
        logExecution('Máquina detenida: Estado final "z" alcanzado.');
        isRunning = false;
        enableControls(true);
    }
}

// Función para ejecutar la máquina automáticamente
async function runMachine() {
    if (isRunning) {
        isRunning = false;
        runMachineButton.textContent = 'Ejecutar';
        enableControls(true);
        return;
    }

    isRunning = true;
    runMachineButton.textContent = 'Detener';
    enableControls(false, true); // Deshabilitar todo excepto Detener

    while (isRunning && currentState !== 'z') {
        step();
        await new Promise(resolve => setTimeout(resolve, 300)); // Pequeña pausa para visualización
    }
    if (currentState === 'z') {
        logExecution('Máquina finalizó su ejecución.');
    }
    isRunning = false;
    runMachineButton.textContent = 'Ejecutar';
    enableControls(true);
}

// Función para habilitar/deshabilitar controles
function enableControls(enable, runningMode = false) {
    loadRulesButton.disabled = !enable;
    resetMachineButton.disabled = !enable;
    stepMachineButton.disabled = !enable;
    inputTapeElement.disabled = !enable;
    rulesFileElement.disabled = !enable;

    if (runningMode) {
        runMachineButton.disabled = false; // Solo el botón de detener está habilitado
    } else {
        runMachineButton.disabled = !enable;
    }
}


// Event Listeners
loadRulesButton.addEventListener('click', loadRules);
resetMachineButton.addEventListener('click', initializeMachine);
stepMachineButton.addEventListener('click', step);
runMachineButton.addEventListener('click', runMachine);

// Inicializar la máquina al cargar la página (sin reglas cargadas aún)
document.addEventListener('DOMContentLoaded', () => {
    initializeMachine();
    enableControls(true); // Asegurarse de que los controles estén habilitados al inicio
});